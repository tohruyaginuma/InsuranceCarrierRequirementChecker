# Insurance Carrier Requirement Checker

A full-stack validation system for insurance quote applications. This system validates applicant data against carrier-specific business requirements before processing quotes.

## Project Structure

```
.
├── backend/                 # Go API server
│   ├── main.go             # Entry point
│   ├── handler/            # HTTP handlers
│   ├── model/              # Domain models
│   ├── store/              # Data store (in-memory)
│   ├── validator/          # Validation logic
│   ├── registry/           # Dependency injection
│   ├── route/              # Routing configuration
│   └── data/
│       └── data.json       # Carrier requirements
│
└── frontend/               # React application
    └── app/
        ├── routes/         # Page components
        └── components/     # UI components
```

## Tech Stack

### Backend
- **Language:** Go 1.19+
- **Framework:** Echo v4
- **Architecture:** Layered (Handler → Store)

### Frontend
- **Framework:** React 19 + React Router 7
- **Styling:** Tailwind CSS 4
- **UI Components:** Radix UI + shadcn/ui
- **State Management:** TanStack Query

## Getting Started

### Prerequisites
- Go 1.19+
- Node.js 20+

### Backend

```bash
cd backend
go run main.go
```

Server starts at `http://localhost:8080`

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Application starts at `http://localhost:5173`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/v1/applicants/` | Create applicant (with validation) |
| GET | `/v1/applicants/` | List all applicants |

### Create Applicant Request

```json
{
  "given_name": "John",
  "surname": "Doe",
  "date_of_birth": "1980-01-01",
  "insurance_status": "CI",
  "prior_carrier": "Allstate",
  "umpd": 5000,
  "collision": null
}
```

### Validation Error Response

```json
{
  "result": "NG",
  "fieldResults": [
    {
      "propertyName": "ApplicantSurname",
      "isValid": false,
      "errorMessage": "Applicant Surname is required."
    }
  ]
}
```

### Success Response

```json
{
  "result": "OK",
  "applicant_id": 1
}
```

