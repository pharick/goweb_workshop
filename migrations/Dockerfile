FROM ubuntu:jammy

# Update system packages
# and install golang-migrate tool
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl && \
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash && \
    apt-get update && \
    apt-get install -y migrate

WORKDIR /usr/var/migrations

# Copy the migration files
COPY ./ ./

# Allow to execute the migration script
RUN chmod +x migrate.sh

# Run the migration
CMD ["./migrate.sh"]
