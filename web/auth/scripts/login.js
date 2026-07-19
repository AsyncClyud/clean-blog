"use strict";

async function SendLoginRequest() {
  const username = document.getElementById("username").value
  const password = document.getElementById("password").value
  const turnstile_token = turnstile.getResponse()

  const response = await fetch("/auth/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Username: username,
      Password: password,
      Turnstile_token: turnstile_token
    })
  })
  if (response.ok) {
    const message = JSON.parse(await response.json())
    if (message != "You has been successfully logined!") {
      turnstile.reset()
      document.getElementById("status").textContent = message
    }
    else {
      document.getElementById("status").textContent = message
      await new Promise(r => setTimeout(r, 2000));
      window.location.replace("/profile")
    }
  }
}
