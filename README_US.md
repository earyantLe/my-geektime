[English](./README_US.md) | [中文](./README.md)


## My Geektime

Not just a downloader, but also an online documentation system

 * Supports one-time caching of Geektime VIP account data for permanent viewing
 * Supports publishing entire courses as online documentation with one click
 * Supports downloading all course audio and video resources to local directories with one click



### [Online Demo](http://8.141.6.243:8090)


### Requirements

- Go 1.25+
- Node.js 20+ (for frontend building)
- Docker & Docker Compose (optional, for containerized deployment)
- MySQL/PostgreSQL/SQLite (database, SQLite by default)


### Installation

#### Method 1: Download Release Version (Recommended)

The simplest and fastest way, directly download pre-compiled binary files from GitHub Releases.

1. Visit [Releases Page](https://github.com/zkep/my-geektime/releases)
2. Download the installation package for your operating system:
   - macOS: `my-geektime-darwin-amd64.tar.gz` or `my-geektime-darwin-arm64.tar.gz`
   - Linux: `my-geektime-linux-amd64.tar.gz`
   - Windows: `my-geektime-windows-amd64.zip`

3. Extract and generate configuration file
```shell
# macOS/Linux
tar -xzf my-geektime-*.tar.gz
./my-geektime cli config
./my-geektime server --config=config_templete.yml

# Windows
# After extracting the zip file
my-geektime.exe cli config
my-geektime.exe server --config=config_templete.yml
```

4. Access the application
Open in browser: http://127.0.0.1:8090




#### Method 2: Docker Compose Deployment

The simplest and fastest deployment method, suitable for production environments.

1. Clone the project
```shell
git clone https://github.com/zkep/my-geektime.git
cd my-geektime/docker
```

2. Start the service
```shell
docker-compose up -d
```

3. Access the application
Open in browser: http://127.0.0.1:8090




#### Method 3: Docker Image Deployment

1. Pull the official image
```shell
# For AMD64 architecture
docker pull --platform=linux/amd64 zkep/mygeektime:latest

# For ARM64 architecture (e.g., Apple Silicon Mac)
docker pull --platform=linux/arm64 zkep/mygeektime:latest
```

2. Prepare configuration file
Create a `config.yml` configuration file (refer to the `config.yml` template in the project root directory)

3. Run the container
```shell
docker run -d \
  --name my-geektime \
  -p 8090:8090 \
  -v $(pwd)/config.yml:/config.yml \
  -v $(pwd)/repo:/repo \
  zkep/mygeektime:latest \
  server --config=/config.yml
```


#### Method 4: Build from Source

Suitable for developers and users who need custom builds.

##### Prerequisites
- Install Go 1.25+
- Install Node.js 20+ and npm

##### Build Steps

1. Clone the project
```shell
git clone https://github.com/zkep/my-geektime.git
cd my-geektime
```

2. Build frontend assets
```shell
make web
```

3. Build backend
```shell
make build
```

4. Prepare configuration file
Copy and modify the configuration template:
```shell
cp config.yml config.local.yml
# Edit config.local.yml as needed
```

5. Run the service
```shell
./my-geektime server --config=config.local.yml
```

6. Access the application
Open in browser: http://127.0.0.1:8090


#### Method 5: Direct Run (Development Mode)

Suitable for quick testing and development debugging.

1. Clone the project and enter the directory
```shell
git clone https://github.com/zkep/my-geektime.git
cd my-geektime
```

2. Install dependencies and run
```shell
# Backend dependencies
go mod download

# Frontend dependencies
cd frontend && npm install && cd ..

# Start development server (frontend)
cd frontend && npm run dev

# Start backend service in another terminal
go run main.go server --config=config.yml
```

3. Access the application
- Frontend development server: http://localhost:3000
- Backend API service: http://127.0.0.1:8090


### Configuration

Key configuration items in `config.yml`:

```yaml
server:
  http_port: 8090          # Service port
  
database:
  driver: sqlite           # Database type: sqlite|mysql|postgres
  source: mygeektime.db    # Database connection string

```

For more configuration options, please refer to the comments in the `config.yml` file.


### Database Support

The project supports multiple databases:

- **SQLite** (default): No additional configuration required, suitable for single-machine deployment
- **MySQL**: Suitable for production environments, better performance
- **PostgreSQL**: Alternative option, powerful features

Example configuration for MySQL:
```yaml
database:
  driver: mysql
  source: root:password@tcp(127.0.0.1:3306)/mygeektime?charset=utf8&parseTime=True&loc=Local
```


### Troubleshooting

1. **Port Already in Use**
   - Modify the `http_port` configuration in `config.yml`
   - Or stop the process using port 8090

2. **Database Connection Failed**
   - Check if the database service is running
   - Verify the connection string in the configuration file
   - Confirm database permission settings

3. **Frontend Assets Loading Failed**
   - Ensure `make web` has been executed to build the frontend
   - Check if the `web/` directory exists and contains build artifacts

4. **Docker Deployment Issues**
   - Confirm Docker and Docker Compose versions
   - Check volume mount path permissions
   - View container logs: `docker-compose logs -f`


### WeChat Sponsor

If you find this useful, your support is appreciated 👏

Sponsor and leave your <b>email account</b> in the message, and I'll send you <b>database courses for reference [Online Demo](http://8.141.6.243:8090)*</b>

<picture>
  <img
    alt="sponsor"
    src="web/sponsor.jpg"
    width="356px"
  />
</picture>


### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.



