"use strict";

async function SendRegisterRequest() {
  const username = document.getElementById("username").value
  const password = document.getElementById("password").value

  const response = await fetch("/auth/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Username: username,
      Password: password,
    })
  })
  if (response.ok) {
    const message = JSON.parse(await response.json())
    document.getElementById("status").textContent = message
    sleep(3500)
    window.location.replace("/")
  }
}

function sleep(time) {
  return new Promise((resolve) => setTimeout(resolve, time));
}
