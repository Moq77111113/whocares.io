# WhoCares.io 🤷‍♂️

> “Today, over 8 million professionals made a great decision: they ignored you.”
> Let’s count them together.

A sarcastically sophisticated web application that tracks professional silence in real-time. Perfect for sharing on LinkedIn when someone drops an unsolicited hot take.

## ✨ Features

- **Dynamic Counter**: Real-time tracking of professionals who chose strategic silence
- **Sarcastic Messaging**: Rotating collection of professionally passive-aggressive messages
- **HTMX Polling**: Live updates every 8 seconds without page refresh
- **Alpine.js Animations**: Smooth transitions for new counts
- **OpenGraph Images**: Dynamic OG images for social sharing
- **Responsive Design**: Professional UI that works on all devices
- **Target Support**: Custom messages via `?target=username` parameter

## 🛠️ Tech Stack

- **Backend**: Go + Echo v4
- **Templates**: Gomponents (type-safe HTML components)
- **Frontend**: HTMX + Alpine.js + Tailwind CSS
- **Images**: Dynamic OG image generation with gg library
- **Architecture**: Clean, `modular Go architecture with service layer

## 📁 Project Structure

```
cmd/server/           # Application entry point
internal/
  ├── components/     # Gomponents templates
  │   ├── layout/     # Base layouts
  │   └── pages/      # Page components
  ├── config/         # Configuration management
  ├── handlers/       # HTTP handlers
  ├── routes/         # Route definitions
  └── services/       # Business logic
      ├── counter.go  # Counter logic
      ├── messages.go # Message generation
      ├── og_image.go # Dynamic OG images
      ├── image.go    # Image management
      └── cron.go     # Cleanup jobs
public/
  ├── css/           # Compiled CSS
  └── og/            # Generated OG images
config.yaml          # Configuration file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/whocares.git
cd whocares

# Install dependencies
go mod tidy

# Compile Tailwind CSS
./tailwindcss -i ./tailwind.css -o public/css/main.css

# Run the application
go run cmd/server/main.go
```

The server starts on `http://localhost:8080`

### Configuration

Create a `config.yaml` file or use environment variables:

```yaml
app:
  title: 'WhoCares.io'
  description: 'Professional Silence Tracker'
  base_url: 'https://whocares.io'

server:
  port: '8080'
  host: 'localhost'

counter:
  base_count: 8000000
  update_seconds: 8
  variation: 500000
```

### CLI Options

```bash
# Custom port
./whocares --port 3000

# Custom host
./whocares --host 0.0.0.0

# Custom config file
./whocares --config /path/to/config.yaml

# Help
./whocares --help
```

## 🧠 Endpoints You Didn't Ask For

- `GET /` – The homepage of glorious indifference.
- `GET /counter` – Live feed of people ignoring you (HTMX-powered).
- `GET /api/og-image` – Generate beautiful silence-sharing images.
- `GET /public/*` – Static files. Because even sarcasm needs CSS.

### Custom Messages

Add a target parameter to personalize the message:

```
https://whocares.io/?target=LinkedInInfluencer
```

## 💬 Example Messages

- "8,392,990 professionals chose silence over your take."
- "8,156,743 industry experts didn't ask for your opinion today."
- "8,284,567 thought leaders successfully avoided your hot takes."
- "8,321,890 decision makers found peace without your input."

## 🏗️ Architecture

This project follows clean Go architecture principles:

- **Separation of Concerns**: Clear boundaries between layers
- **Type-Safe Templates**: Gomponents for compile-time HTML safety
- **HTMX Integration**: Modern web interactivity without JavaScript complexity
- **Configuration Management**: Cobra + Viper for professional CLI and config
- **Service Layer**: Business logic separated from HTTP concerns

## 🤝 Contributing

We welcome contributions! Please follow these guidelines:

1. **Code Style**: Follow Go conventions and keep functions under 90 lines
2. **Sarcasm Level**: Maintain professional passive-aggressiveness
3. **Testing**: Add tests for new features
4. **Documentation**: Update README for significant changes

### Development Workflow

```bash
# Create feature branch
git checkout -b feature/your-feature

# Make changes
# ...

# Run tests
go test ./...

# Build and test
go build ./cmd/server
./server

# Commit and push
git commit -m "feat: add your feature"
git push origin feature/your-feature
```

## 📝 License

MIT License - See [LICENSE](LICENSE) file for details.

## 🎯 Philosophy

> "Sometimes the world needs to know exactly how many people didn't ask."

WhoCares.io exists in the professional space between LinkedIn thought leadership and Slack channel expertise. It's for those moments when someone shares their "innovative insights" and you want to respond with data-driven professional silence metrics.

## � Star History

If this project helped you maintain professional composure in the face of unsolicited opinions, consider giving it a star! ⭐

---

**Built with 💼 by professionals who didn't ask for your feedback either.**
