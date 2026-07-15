"use strict";

function GetPathValue() {
  var URL = (window.location.pathname)
  var LastIndexURL = URL.substring(0, URL.length - 1).lastIndexOf("/")
  const Id = URL.substring(LastIndexURL+1, URL.length+1)
  return Id
}
async function Fetch_Article() {
  const Id = GetPathValue()
  const article_request = await fetch(`/api/articles/${Id}`, {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (article_request.ok) {
    const article = JSON.parse(await article_request.json());
    const main_element = document.getElementById("main");
    const article_element = document.createElement("div");
    article_element.setAttribute("class", "articles");
    article_element.setAttribute("id", "article");
    article_element.innerHTML = `
    <h3 id = "title">${article.Title}</h3>
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

  const article_author_request = await fetch("/api/users", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Author: Number(author_id)
      })
  })
  if (article_author_request.ok) {
    const article_id = GetPathValue()
    const article_element = document.getElementById("article")
    const delete_button = document.createElement("button")
    const update_button = document.createElement("button")

    delete_button.textContent = `Delete article`
    delete_button.setAttribute("onclick", `SendDeleteRequest()`)
    update_button.textContent = `Update article`
    update_button.setAttribute("onclick", `UpdatePageRedirect(${article_id})`)

    article_element.appendChild(delete_button)
    article_element.appendChild(update_button)
  }
}

async function SendCreateRequest() {
  const title = document.getElementById("title").value
  const content = document.getElementById("content").value

  const create_request = await fetch("/api/articles", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Title: title,
      Content: content
      })
  })
  if (create_request.ok) {
    const message = JSON.parse(await create_request.json())
    document.getElementById("status").textContent = message
  }
}

async function SendDeleteRequest() {
  const article_id = GetPathValue()
  const author_id = document.getElementById("author_id").textContent

  const delete_request = await fetch("/api/articles", {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Id: Number(article_id),
      Author: Number(author_id)
    })
  })
  if (delete_request.ok) {
   document.location.replace("/")
  }
}

async function SendUpdateRequest() {
  const new_title = document.getElementById("new_title").value
  const new_content = document.getElementById("new_content").value
  const author_id = document.getElementById("author_id").textContent

  const update_request = await fetch("/api/articles", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Id: Number(GetPathValue()),
      Title: new_title,
      Content: new_content,
      Author: Number(author_id)
    })
  })
  if (update_request.ok) {
    const message = JSON.parse(await update_request.json())
    document.getElementById("status").textContent = message
 }
}

function UpdatePageRedirect(article_id) {
  window.location.replace(`/article/update/${article_id}`)
}
