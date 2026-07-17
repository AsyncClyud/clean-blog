"use strict";

async function SendRegisterRequest() {
  const username = document.getElementById("username").value
  const password = document.getElementById("password").value
  const turnstile_token = turnstile.getResponse()

  const response = await fetch("/auth/register", {
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
    if (message != "Account has been created!") {
      turnstile.reset()
      document.getElementById("status").textContent = message
    }
    document.getElementById("status").textContent = message
    await new Promise(r => setTimeout(r, 2000));
    window.location.replace("/profile")
  }
}
