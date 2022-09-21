# GH Keys

Use github public keys to authorize ssh connections.

### Install

```Bash

go install github.com/rtgnx/ghkeys/cmd@latest

```

### SSHD Config

**Allow any valid user from github**

```
AuthorizedKeysCommand /usr/bin/ghkeys %u 2> /tmp/ghkey.log
AuthorizedKeysUser nobody
```

**Allow specific user**

```
AuthorizedKeysCommand /usr/bin/ghkeys --users rtgnx,myfriend %u 2> /tmp/ghkey.log
AuthorizedKeysUser nobody
```

Users will be looked up against /etc/passwd and check if user id is > 1000





