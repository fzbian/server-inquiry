# Server Inquiry

This code is designed to perform command line requests on a server from an external application via an API. Security is based on the creation of a token each time the application is started. This will help communication between applications and servers in a fast and secure way. By generating a unique token at each login, it ensures that only the authorized application can send requests to the server and receive information back. In addition, by using an API, requests can be sent and received in a standardized format, which facilitates integration between different systems and applications.

## How it works

It works by creating a server by opening the necessary ports to receive requests. Before any request can be made, the server validates that the token sent is valid. Once authenticated, requests can be sent to the server by specifying the necessary parameters, such as the operating system and the command to be executed. The server will receive the request, verify the parameters and, if valid, execute the specified command on the chosen operating system. Once executed, the server returns the result of the execution to the client. This process is automated, allowing actions to be performed on the server quickly and securely via HTTP requests.

![](http://imgfz.com/i/43myVQC.jpeg)

**Generating the token**

After running the start command the application will generate a local token, which will only work when the application is running, when the application closes the token will not be valid for anything, when the application runs again it will generate a different one.

![](https://media.discordapp.net/attachments/1052025865803939880/1064676411639603220/image.png?width=1020&height=52)

 - Health
   - Authorized 
      ![](https://media.discordapp.net/attachments/1052025865803939880/1064676912510804030/image.png)
   - Unauthorized
      ![](https://media.discordapp.net/attachments/1052025865803939880/1064677229440806932/image.png)

 - Command
   - Windows (PowerShell)
     - `systeminfo`
       ![](https://media.discordapp.net/attachments/1052025865803939880/1064678216398278776/image.png?width=940&height=480)
     - `echo "Hello Windows!"`
       ![](https://media.discordapp.net/attachments/1052025865803939880/1064678336749637713/image.png?width=1020&height=215)
   - Linux (Ubuntu)
       - `echo "Hello world"`
         ![](https://media.discordapp.net/attachments/1052025865803939880/1064684930346516559/image.png?width=1020&height=219)
       - `cat /proc/cpuinfo`
         ![](https://media.discordapp.net/attachments/1052025865803939880/1064685542903652433/image.png?width=922&height=480)