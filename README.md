# Minecraft Server Java Edition

<p>
  <a aria-label="Deploy on Host Factor" href="https://hostfactor.io/games/minecraft">
    <img src="https://img.shields.io/badge/Deploy-Host%20Factor-%234f6ac6?labelColor=1b1c1d&style=for-the-badge" alt="">
  </a>
  <a aria-label="Build status" href="https://github.com/hostfactor/minecraft-server/actions/workflows/build_latest.yml">
    <img src="https://img.shields.io/github/workflow/status/hostfactor/minecraft-server/Build%20latest?style=for-the-badge&labelColor=1b1c1d" alt="">
  </a>
</p>


A lightweight Docker image tagged for all Minecraft versions containing only the server and minimal dependencies to run
it.

## Usage

All images are tagged and packed with a specific Minecraft version. For example, to run the default Java version
for `1.18.1`:

### Run the latest

```
docker run -p 25565:25565 ghcr.io/hostfactor/minecraft-server
```

### Run `1.18.1`

```
docker run -p 25565:25565 ghcr.io/hostfactor/minecraft-server:1.18.1
```

### Run with different Java versions

If you want to use a specific Java version e.g. `Java 11` for version `1.12`, you can run

```
docker run -p 25565:25565 ghcr.io/hostfactor/minecraft-server:1.12-java-11
```

## Java versions

Every default tag for a version (e.g. `1.18`) will run the latest Java version that will work for that version. The
supported Java versions are:

| Java Version | Supported Version | Default Version     |
|--------------|-------------------|---------------------|
| 17           | All               | \>= `1.18`          |
| 16           | < `1.18`          | < `1.18`, >= `1.17` |
| 11           | < `1.17`          | < `1.17`, >= `1.12` |
| 8            | < `1.12`          | < `1.12`            |

## Advanced examples

### Existing world

In order to use an existing world, simply copy the folder containing your Minecraft world into the `/server/world`
folder in the image.

```
docker run -p 25565:25565 -v /path/to/world:/server/world ghcr.io/hostfactor/minecraft-server
```

**Note**: Your world folder must contain the `level.dat` file in the root.

### Custom `server.properties`

Minecraft servers are configured via
a [server.properties](https://minecraft.fandom.com/wiki/Server.properties#Java_Edition_3) file. If you want to use
custom options for this file, use the following command.

```
docker run -p 25565:25565 -v /path/to/server.properties:/server/server.properties ghcr.io/hostfactor/minecraft-server
```

Where `/path/to/server.properties` is the absolute path to the `server.properties` file on your computer.

### Args

You can run your Minecraft server
with [args](https://minecraft.fandom.com/wiki/Tutorials/Setting_up_a_server#Minecraft_options) by utilizing the `OPTS`
env var e.g.

```
$ docker run -p 25565:25565 -e OPTS="--help" ghcr.io/hostfactor/minecraft-server
Option                   Description                                         
------                   -----------                                         
--bonusChest                                                                 
--demo                                                                       
--eraseCache                                                                 
--forceUpgrade                                                               
--help                                                                       
--initSettings           Initializes 'server.properties' and 'eula.txt', then
                           quits                                             
--nogui                                                                      
--port <Integer>         (default: -1)                                       
--safeMode               Loads level with vanilla datapack only              
--serverId <String>                                                          
--singleplayer <String>                                                      
--universe <String>      (default: .)                                        
--world <String> 
```

### Java options

You can configure Java options (most importantly JVM heap size) via the `_JAVA_OPTIONS` env var e.g.

```
docker run -p 25565:25565 -e _JAVA_OPTIONS="-Xmx4G -Xms4G" ghcr.io/hostfactor/minecraft-server 
```

The above will allocate 4GB of memory, no more, no less.
See [here](https://minecraft.fandom.com/wiki/Tutorials/Setting_up_a_server#Java_options) for more info.
