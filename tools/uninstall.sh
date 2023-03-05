#!/bin/bash

# Removes the ServerInquiry binary from /usr/local/bin
sudo rm /usr/local/bin/server-inquiry
# Removes the PATH line in .profile
sudo sed -i '/# ServerInquiry/d' ~/.profile
sudo sed -i '/export PATH=\$PATH:\/usr\/local\/bin/d' ~/.profile
# Deletes the PATH line in the current session
export PATH=$(echo $PATH | sed -e 's/:\/usr\/local\/bin//g')
# Uninstallation completed message
echo "ServerInquiry (1.0.0b) has been uninstalled from your machine."