package ghkeys

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/user"
	"regexp"
	"strconv"
	"strings"
)

const NameRegex = `^[a-zA-Z0-9-]+$`

type KeySource func(string) ([]string, error)

var Sources = map[string]KeySource{
	"github": Github,
}

func Keys(name string, cfg Config) (keys []string, err error) {
	if !regexp.MustCompile(NameRegex).MatchString(name) {
		return []string{}, fmt.Errorf("username %s does not match %s", name, NameRegex)
	}

	if len(cfg.AllowedUsers) != 0 && !Has(name, cfg.AllowedUsers) {
		return []string{}, fmt.Errorf("%s is not allowed", name)
	}

	u, err := user.Lookup(name)
	if err != nil {
		return []string{}, fmt.Errorf("user not present in /etc/passwd")
	}

	if uid, err := strconv.ParseUint(u.Uid, 10, 32); err != nil || uid < 1000 {
		return []string{}, fmt.Errorf("only uid > 1000 are allowed, got %d", uid)
	}

	if len(cfg.AllowedSources) == 0 {
		for _, p := range Sources {
			if k, err := p(name); err == nil {
				keys = append(keys, k...)
			}
		}
	} else {
		for _, allowed := range cfg.AllowedSources {
			if p, ok := Sources[allowed]; ok {
				if k, err := p(name); err == nil {
					keys = append(keys, k...)
				}
			}
		}
	}

	return keys, err
}

func Has[T comparable](k T, slice []T) bool {
	for _, v := range slice {
		if v == k {
			return true
		}
	}
	return false
}

func Github(username string) ([]string, error) {
	res, err := http.Get(fmt.Sprintf("https://github.com/%s.keys", username))
	if err != nil {
		return []string{}, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(b), "\n"), nil
}

func Local(pattern string) KeySource {
	return func(s string) ([]string, error) {
		// Remove path traversal attempts
		s = strings.ReplaceAll(s, "/", "")

		b, err := ioutil.ReadFile(fmt.Sprintf(pattern, s))

		if err != nil {
			return []string{}, err
		}

		return strings.Split(string(b), "\n"), err
	}
}
