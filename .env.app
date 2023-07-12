APP_ENV=dev 

# Server settings:
SERVER_URL="0.0.0.0:3000"
SERVER_READ_TIMEOUT=60
APP_VERSION="1.0.0"
# JWT settings:
JWT_SECRET_KEY="fa1fb6c834c627461a71809c25a1f10a"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=300

# Database settings:
DB_SERVER_URL="host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_LIFETIME_CONNECTIONS=2