#!/bin/sh

# Create Redis config directory
mkdir -p /usr/local/etc/redis

# Allow Redis to accept connections from any IP address
echo "bind 0.0.0.0" > /usr/local/etc/redis/redis.conf

# Add password protection
if [ -n "$REDIS_PASSWORD" ]; then
  echo "requirepass $REDIS_PASSWORD" >> /usr/local/etc/redis/redis.conf
fi

# Start Redis with config
exec redis-server /usr/local/etc/redis/redis.conf