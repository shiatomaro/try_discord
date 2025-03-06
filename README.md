# GoTryBot - A Fun Discord Bot in Go

GoTryBot is a simple Discord bot written in **Go** that responds to commands with jokes and fun messages. It uses the [DiscordGo](https://github.com/bwmarrin/discordgo) library to interact with Discord and supports basic command handling.

## Features  
- Responds to the command `!gotrybot gohappy` with a random tech-related joke  
- Lists available commands using `!gotrybot Commands`  
- Loads bot token securely from a `.env` file  
- Uses embedded messages for better formatting  

## Prerequisites  
- Go installed on your system  
- A Discord bot token (set in a `.env` file)  

## Installation & Setup  

1. Clone the repository:  
   ```sh
   git clone https://github.com/your-username/gotrybot.git
   cd gotrybot
   ```  
2. Install dependencies:  
   ```sh
   go mod tidy
   ```  
3. Create a `.env` file and add your bot token:  
   ```ini
   DISCORD_BOT_TOKEN=your_discord_bot_token
   ```  
4. Run the bot:  
   ```sh
   go run main.go
   ```  

## Usage  
Once the bot is running and invited to a server, use the following commands:  
- `!gotrybot Commands` → Displays available commands  
- `!gotrybot gohappy` → Responds with a random joke  

## License  
This project is licensed under the MIT License.  

