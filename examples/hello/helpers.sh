#!/usr/bin/env bash
set -euo pipefail

install() {
  echo Install Hello
}

upgrade() {
  echo World 2.0
}

uninstall() {
  echo Goodbye World
}

create-databrickscfg() {
  cat <<EOF >> ~/.databrickscfg
[DEFAULT]
host = $1
token = $2
EOF

cat ~/.databrickscfg
}

# Call the requested function and pass the arguments as-is
"$@"
