"use strict";

async function Logout() {
  const response = await fetch("/api/logout", {
    method: "POST",
    credentials: "include"
  })
  if (response.ok) {
    window.location.replace("/")
  }
}
