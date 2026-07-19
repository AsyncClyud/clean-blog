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
      const comment_creator_element = document.getElementById("comment_creator")

      const profile = document.createElement("a")
      profile.href = `/profile`
      profile.textContent = `Profile`
      profile.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")
      const create_article = document.createElement("a")
      create_article.href = `/article/create`
      create_article.textContent = `Create article`
      create_article.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")
      const logout = document.createElement("button")
      logout.setAttribute("onclick", "Logout()")
      logout.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")
      logout.textContent = `Logout`

      header_element.appendChild(profile)
      header_element.append(create_article)
      header_element.append(logout)

      comment_creator_element.innerHTML = `
        <h2 class="text-[4vh] text-center font-[JetBrains_Mono] m-[10px]">Leave a comment</h2>
        <textarea class="w-[35vw] bg-[white] text-[black] resize-none rounded-[5px] m-[5px]" id="comment_content" rows="10" placeholder="Share your thoughts..."></textarea> <br>
        <button class="w-fit bg-[white] text-[black] text-[JetBrains_Mono] rounded-[3px] m-[5px] p-[5px]" type="button" onclick="SendCreateCommentRequest()">Post comment</button>
        <p id="status"></p>
        <div class="cf-turnstile" id="turnstile-widget" data-sitekey="0x4AAAAAAD2voHPreG9maJ8u" data-theme="dark"></div>
        `
    }
    else {
      const header_element = document.getElementById("header")

      const registration = document.createElement("a")
      registration.href = `/auth/register`
      registration.textContent = `Registration`
      registration.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")

      const login = document.createElement("a")
      login.href = `/auth/login`
      login.textContent = `Login`
      login.setAttribute("class", "bg-[#2a323c] text-[gainsboro] rounded-[3px] p-[5px] m-[5px]")

      header_element.appendChild(registration)
      header_element.appendChild(login)
    }
  }
}

export default Auth()
