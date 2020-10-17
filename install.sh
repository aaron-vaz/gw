#!/bin/bash -e
{
  has_command() {
    if ! command -v "$1" >/dev/null 2>&1; then
      echo 1
    else
      echo 0
    fi
  }

  GW_HOME=${GW_HOME}
  GW_VERSION=${GW_VERSION:-latest}

  if [ "$GW_HOME" == "" ] || [ "$GW_HOME" == "$HOME/.gw" ]; then
    GW_HOME=$HOME/.gw
  fi

  # curl looks for HTTPS_PROXY while wget for https_proxy
  https_proxy=${https_proxy:-$HTTPS_PROXY}
  HTTPS_PROXY=${HTTPS_PROXY:-$https_proxy}

  if [ "$GW_GET" == "" ]; then
    if [ 0 -eq $(has_command curl) ]; then
      GW_GET="curl -sL"
    elif [ 0 -eq $(has_command wget) ]; then
      GW_GET="wget -qO-"
    else
      echo "[ERROR] This script needs wget or curl to be installed."
      exit 1
    fi
  fi

  if [ "$GW_VERSION" == "latest" ]; then
    GW_VERSION=$($GW_GET https://github.com/aaron-vaz/gw/raw/master/version)
  fi

  case "$OSTYPE" in
  darwin*)
    BINARY_URL=https://github.com/aaron-vaz/gw/releases/download/${GW_VERSION}/gw_darwin_amd64
    ;;
  linux*)
    BINARY_URL=https://github.com/aaron-vaz/gw/releases/download/${GW_VERSION}/gw_linux_amd64
    ;;
  *)
    echo "Unsupported OS $OSTYPE. If you believe this is an error - please create a ticket at https://github.com/aaron-vaz/gw/issues."
    exit 1
    ;;
  esac

  echo "Installing gw v$GW_VERSION..."
  echo

  mkdir -p ${GW_HOME}/bin

  $GW_GET ${BINARY_URL} >${GW_HOME}/bin/gw && chmod a+x ${GW_HOME}/bin/gw

  echo "Installation finished"
  echo

  echo "Please add the following to a location where it can be sourced"
  echo
  echo "if [ -f ${GW_HOME}/bin/gw ]; then"
  echo "  export PATH=${GW_HOME}/bin:\$PATH"
  echo "fi"
}
