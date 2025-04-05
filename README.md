# Push Run App

A Strava-like application for the Russian market.

## Features

- User authentication with GitHub OAuth
- Activity tracking and statistics
- Social features (friends, likes, comments)
- Achievements system
- Challenges
- Route planning with Yandex Maps integration
- Weather information

## Tech Stack

- Frontend: Vue.js, TypeScript
- Backend: NestJS, TypeScript
- Database: PostgreSQL
- Authentication: JWT, GitHub OAuth
- Maps: Yandex Maps API

## Prerequisites

- Node.js (v16 or later)
- PostgreSQL
- GitHub OAuth App credentials
- Yandex Maps API key

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/push-run-app.git
cd push-run-app
```

2. Install frontend dependencies:
```bash
cd frontend
npm install
```

3. Install backend dependencies:
```bash
cd ../backend-nest
npm install
```

4. Create a `.env` file in the backend-nest directory with the following variables:
```env
# Server
PORT=3000

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_DATABASE=push_run_app

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRES_IN=7d

# GitHub OAuth
GITHUB_CLIENT_ID=your-github-client-id
GITHUB_CLIENT_SECRET=your-github-client-secret
GITHUB_CALLBACK_URL=http://localhost:3000/auth/github/callback

# Yandex Maps
YANDEX_MAPS_API_KEY=your-yandex-maps-api-key
```

5. Start the development servers:
```bash
# Terminal 1 - Backend
cd backend-nest
npm run start:dev

# Terminal 2 - Frontend
cd frontend
npm run dev
```

The application will be available at:
- Frontend: http://localhost:5173
- Backend: http://localhost:3000
- API Documentation: http://localhost:3000/api

## Project Structure

```
push-run-app/
├── frontend/            # Vue.js frontend application
│   ├── src/
│   │   ├── components/  # Vue components
│   │   ├── views/       # Page components
│   │   ├── router/      # Vue Router configuration
│   │   ├── store/       # Vuex store
│   │   └── assets/      # Static assets
│   └── package.json     # Frontend dependencies
│
├── backend-nest/        # NestJS backend application
│   ├── src/
│   │   ├── models/      # Data models and DTOs
│   │   ├── modules/     # Feature modules
│   │   ├── config/      # Configuration
│   │   └── main.ts      # Application entry point
│   └── package.json     # Backend dependencies
│
└── README.md            # Project documentation
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
