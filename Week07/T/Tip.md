shell

```shell
# shell exec


# SYSTEM ENHANCED SETTINGS

# Same as -e
set -o errexit

# Exit  immediately  if a pipeline (which may consist of a single simple command),  a subshell command enclosed in paren‐theses, or one of the commands executed as part of a command list enclosed by braces (see SHELL  GRAMMAR  above)  exits with a non-zero status.  The shell does not exit if the command that fails is part of the command list immediately fol‐lowing a while or until keyword, part of the test following the if or elif reserved words, part of any command executed in  a  && or || list except the command following the final && or ||, any command in a pipeline but the last, or if the command's return value is being inverted with !.  A trap on ERR, if set, is executed  before  the  shell  exits.   This option  applies  to  the  shell environment and each subshell environment separately (see COMMAND EXECUTION ENVIRONMENT above), and may cause subshells to exit before executing all the commands in the subshell
set -e


# Same as -u
set -o nounset

# Treat unset variables and parameters other than the special parameters "@" and "*" as an error when performing  parame‐ter expansion.  If expansion is attempted on an unset variable or parameter, the shell prints an error message, and, if not interactive, exits with a non-zero status
set -u


# If set, the return value of a pipeline is the value of the last (rightmost) command to  exit  with  a  non-zero status, or zero if all commands in the pipeline exit successfully.  This option is disabled by default.
set -o pipefail


```

