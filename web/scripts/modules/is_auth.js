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

      header_element.appendChild(profile)
    }
    else {
      const header_element = document.getElementById("header")

      const registration = document.createElement("a")
      registration.href = `/auth/register`
      registration.textContent = `Registration`
      header_element.appendChild(registration)

      const separator = document.createElement("br")

      const login = document.createElement("a")
      login.href = `/auth/login`
      login.textContent = `Login`

      header_element.appendChild(registration)
      header_element.appendChild(login)
    }
  }
}

export default Auth()
