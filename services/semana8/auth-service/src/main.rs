use axum::{routing::get, routing::post, Router, Json};
use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
struct LoginRequest {
    username: String,
    password: String,
}

#[derive(Serialize)]
struct LoginResponse {
    token: String,
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/health", get(health))
        .route("/login", post(login));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8081").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn health() -> &'static str {
    "ok"
}

async fn login(Json(body): Json<LoginRequest>) -> Json<LoginResponse> {
    // Aquí iría la lógica de autenticación, por ahora simplemente devolvemos un token de ejemplo
    let token = format!("token_for_{}", body.username);
    Json(LoginResponse { token })
}