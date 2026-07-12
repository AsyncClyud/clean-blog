"use strict";

async function FetchUserProfile() {
  const response = await fetch("/api/profile", {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (response.ok) {
    const data = JSON.parse(await response.json())
    var main_element = document.getElementById("main")
    var user_info = document.createElement("div")
    user_info.setAttribute("id", "profile")
    user_info.innerHTML = `
      <img src="profile/images/bitmap.png" alt="Avatar" width="200" height="200">
      <h3>${data.Username}</h3>
      <p>${data.Bio}</p>
      <p>${data.Created_at}</p>
    `
    main_element.appendChild(user_info)
    document.title = `${data.Username} profile`
  }
  if (response.status == 401) {
    window.location.replace("/auth/register")
  }
}
