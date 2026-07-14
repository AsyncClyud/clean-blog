"use strict";

function GetPathValue() {
  var URL = (window.location.pathname)
  var LastIndexURL = URL.substring(0, URL.length - 1).lastIndexOf("/")
  const Id = URL.substring(LastIndexURL+1, URL.length+1)
  return Id
}
async function Fetch_Article() {
  const Id = GetPathValue()
  const response = await fetch(`/api/articles/${Id}`, {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (response.ok) {
    const article = JSON.parse(await response.json());
    const main_element = document.getElementById("main");
    const article_element = document.createElement("div");
    article_element.setAttribute("class", "article");
    article_element.setAttribute("id", Id);
    article_element.innerHTML = `
    <h3 id = "title">${article.Title}<h3>
    <p id ="content">${article.Content}</p>
    <p id="created_at">${article.Created_at}</p>
    <p>Article Author ID:</p>
    <p id="author_id">${article.Author}</p>
    `;
    main_element.appendChild(article_element);
    document.title = article.Title

    await GetArticleAuthor()
  }
}

async function GetArticleAuthor() {
  const author_id = document.getElementById("author_id").textContent

  const response = await fetch("/api/users", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Author: Number(author_id)
      })
  })
  if (response.ok) {
    const main_element = document.getElementById("main")
    const delete_button = document.createElement("button")
//  const update_button = document.createElement("button")

    delete_button.textContent = `Delete article`
    delete_button.setAttribute("onclick", "SendDeleteRequest()")
//  update_button.textContent = `Update article`

    main_element.appendChild(delete_button)
//  main_element.appendChild(update_button)
  }
}

async function SendCreateRequest() {
  const title = document.getElementById("title").value
  const content = document.getElementById("content").value

  const response = await fetch("/api/articles", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Title: title,
      Content: content
      })
  })
  if (response.ok) {
    const message = JSON.parse(await response.json())
    document.getElementById("status").textContent = message
  }
}

async function SendDeleteRequest() {
  const article_id = GetPathValue()
  const author_id = document.getElementById("author_id").textContent

  const response = await fetch("/api/articles", {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Id: Number(article_id),
      Author: Number(author_id)
    })
  })
  if (response.ok) {
   document.location.replace("/")
  }
}
