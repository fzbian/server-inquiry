#!/bin/bash

# Download the ServerInquiry binary
sudo curl -LO https://github.com/fzbian/server-inquiry/releases/download/v1.0.0b/server-inquiry
# Assign execution permissions to the binary
sudo chmod +x server-inquiry
# Move the binary to /usr/local/bin
sudo mv server-inquiry /usr/local/bin/
# Add the binary to the PATH in .profile
echo -e "\n# ServerInquiry" >> ~/.profile
echo "export PATH=\$PATH:/usr/local/bin" >> ~/.profile
# Adds the binary to the PATH in the current session
export PATH=$PATH:/usr/local/bin
# Installation completed message
echo "ServerInquiry (1.0.0b) is ready for use."