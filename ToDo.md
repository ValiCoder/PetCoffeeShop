Coffee Shop Application: Feature Road

Containerized Environment: Multi-container setup using Docker Compose (Go application + PostgreSQL database).

Database Integration: Fully configured PostgreSQL database running on a custom port (5444) to avoid host system conflicts.

Strict Environment Management: Configuration isolation using .env files for secure database credential management (DB_USER, DB_PASSWORD, DB_NAME).

Database Connectivity: Established a persistent, secure IDE-to-Database connection within GoLand for schema introspection and querying.

Menu Catalog API: Backend endpoints to fetch and stream coffee items and desserts from the database.

Interactive Frontend UI: Client-side interface to display the coffee shop menu dynamically using responsive item cards (tiles).

Basic Order Processing: Initial endpoints to accept incoming orders from clients and store them in the system.

Multi-Role System: Implementation of distinct user roles to secure the platform:

User (Client): Can browse the menu and place orders.

Barista: Can view active orders and update their fulfillment status.

Admin: Can manage menu items (add/edit/delete).

Superuser (Owner): Full admin access + staff management and financial reporting.

JWT Authentication: Secure token-based login and signup system to protect restricted backend API routes.

"Add Item" UI Tile: A specialized white tile with a large + icon instead of an image, visible exclusively to users logged in as Admin or Superuser.

Dynamic Item Creator Interface: A frontend form triggered by the + tile, containing the following validation fields:

Name (String)

Category (Dropdown: Coffee, Tea, Desserts, etc.)

Price (Positive Numeric)

Volume/Size (String/Numeric: e.g., 250ml, 350ml, 450ml)

Description/Ingredients (Text)

Image Upload (File/URL string)

Menu Mutator Endpoints: Secure backend API routes (POST, PUT, DELETE) protected by admin-role middleware.

API Documentation (Swagger)
Automated OpenAPI Spec: Integration of swaggo/swag into the Go backend.

Interactive API Dashboard: A locally hosted web UI (/swagger/index.html) where i can live-test endpoints using a "Try it out" feature without needing external tools like Postman.