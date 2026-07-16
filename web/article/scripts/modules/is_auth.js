"use strict";

async function Auth() {
  const response = await fetch("/api/auth", {
    method: "GET",
    headers: { Accept: "application/json" }
  })
  if (response.ok) {
    const data = await response.json()
    if (data.authorized == true) {
      const header_element = document.getElementById("header")

      const profile = document.createElement("a")
      profile.href = `/profile`
      profile.textContent = `Profile`
      const create_article = document.createElement("a")
      create_article.href = `/article/create`
      create_article.textContent = `Create article`
      const logout = document.createElement("button")
      logout.setAttribute("onclick", "Logout()")
      logout.textContent = `Logout`

      header_element.appendChild(profile)
      header_element.append(create_article)
      header_element.append(logout)
    }
    else {
      const header_element = document.getElementById("header")

      const registration = document.createElement("a")
      registration.href = `/auth/register`
      registration.textContent = `Registration`
      header_element.appendChild(registration)

      const login = document.createElement("a")
      login.href = `/auth/login`
      login.textContent = `Login`

      header_element.appendChild(registration)
      header_element.appendChild(login)
    }
  }
}

export default Auth()
