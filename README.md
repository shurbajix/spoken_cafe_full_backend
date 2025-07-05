# Spoken Cafe Backend

A modern, scalable backend architecture for the Spoken Cafe application.

## 🏗️ Architecture Overview

```
backend/
├── src/
│   ├── config/           # Configuration files
│   ├── controllers/      # Request handlers
│   ├── middleware/       # Custom middleware
│   ├── models/          # Data models and schemas
│   ├── routes/          # API route definitions
│   ├── services/        # Business logic
│   ├── utils/           # Utility functions
│   └── types/           # TypeScript type definitions
├── tests/               # Test files
├── docs/                # API documentation
├── scripts/             # Build and deployment scripts
└── docker/              # Docker configuration
```

## 🚀 Quick Start

### Prerequisites
- Node.js 18+
- npm or yarn
- Firebase CLI
- Docker (optional)

### Installation
```bash
cd backend
npm install
```

### Environment Setup
```bash
cp .env.example .env
# Edit .env with your configuration
```

### Development
```bash
npm run dev
```

### Production Build
```bash
npm run build
npm start
```

## 📁 Directory Structure

### `/src/config`
- Database configuration
- Firebase configuration
- Environment variables
- App settings

### `/src/controllers`
- HTTP request handlers
- Input validation
- Response formatting

### `/src/middleware`
- Authentication middleware
- Error handling
- Request logging
- CORS configuration

### `/src/models`
- Data models and schemas
- Database models
- Type definitions

### `/src/routes`
- API route definitions
- Route grouping
- Route middleware

### `/src/services`
- Business logic
- External API integrations
- Payment processing
- Notification services

### `/src/utils`
- Helper functions
- Common utilities
- Constants

### `/src/types`
- TypeScript interfaces
- Type definitions
- API response types

## 🔧 Features

- **Authentication & Authorization**: JWT-based auth with role-based access
- **Payment Processing**: Integration with multiple payment gateways
- **Real-time Notifications**: WebSocket support for live updates
- **File Upload**: Secure file handling with cloud storage
- **API Documentation**: Auto-generated OpenAPI/Swagger docs
- **Error Handling**: Comprehensive error management
- **Logging**: Structured logging with different levels
- **Testing**: Unit and integration tests
- **Docker Support**: Containerized deployment
- **CI/CD**: Automated testing and deployment

## 🔐 Security Features

- Input validation and sanitization
- Rate limiting
- CORS protection
- Helmet.js security headers
- JWT token management
- Password hashing
- SQL injection prevention

## 📊 Database Schema

The backend supports multiple database types:
- **Firestore**: Primary NoSQL database
- **PostgreSQL**: Relational data (optional)
- **Redis**: Caching and sessions

## 🚀 Deployment

### Firebase Functions
```bash
firebase deploy --only functions
```

### Docker
```bash
docker build -t spoken-cafe-backend .
docker run -p 3000:3000 spoken-cafe-backend
```

### Environment Variables
See `.env.example` for required environment variables.

## 📝 API Documentation

API documentation is available at `/api/docs` when running in development mode.

## 🧪 Testing

```bash
# Run all tests
npm test

# Run tests in watch mode
npm run test:watch

# Run tests with coverage
npm run test:coverage
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License. 