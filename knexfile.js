
// Update with your config settings.

require('dotenv').config()

module.exports = {

  development: {
    client: 'sqlite3',
    connection: {
      filename: './dev.db',
    },
    migrations: {
      directory: './database/migrations',
    },
    seeds: {
      directory: './database/seeds',
    }
  },
  production: {
    client: 'sqlite3',
    connection: {
      filename: './web_monitor.db',
    },
    migrations: {
      directory: './database/migrations',
    },
    seeds: {
      directory: './database/seeds',
    }
  }
};
