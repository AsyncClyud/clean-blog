"use strict";

function GetPathValue() {
  var URL = (window.location.pathname)
  var LastIndexURL = URL.substring(0, URL.length - 1).lastIndexOf("/")
  const Id = URL.substring(LastIndexURL+1, URL.length+1)
  return Id
}
async function FetchArticle() {
  const Id = GetPathValue()
  const get_article_request = await fetch(`/api/articles/${Id}`, {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (get_article_request.ok) {
    const article = JSON.parse(await get_article_request.json());
    const article_element = document.getElementById("article");
    article_element.setAttribute("class", "articles");
    article_element.innerHTML = `
    <h3 id = "title">${article.Title}</h3>
    <p id ="content">${article.Content}</p>
    <p class="data" id="created_at">${article.Created_at}</p>
    <p>Article Author ID:</p>
    <p id="author_id">${article.Author}</p>
    <h3 class="comments_title">Comments</h3>
    <div class="comments" id="comments"></div>
    `;
    document.title = article.Title

    await GetArticleAuthor()
  }
  if (document.getElementById("author_id").textContent == "0") {
    window.location.replace("/not_found")
  }
}

async function GetArticleAuthor() {
  const id = GetPathValue()
  const author_id = document.getElementById("author_id").textContent

  const article_author_request = await fetch("/api/users", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Author: Number(author_id)
      })
  })
  if (article_author_request.ok) {
    const main_element = document.getElementById("main")
    const actions_element = document.createElement("div")
    actions_element.setAttribute("class", "actions")
    actions_element.innerHTML = `
      <button type="button" onclick="SendDeleteArticleRequest()">Delete Article</button>
      <button type="button" onclick="UpdatePageRedirect(${id})">Update Article</button>
      `
    main_element.appendChild(actions_element)
  }
}

async function FetchArticleComments() {
  const Id = GetPathValue()

    const comments_request = await fetch(`/api/comments/${Id}`, {
    method: "GET",
    headers: { "Accept": "application/json" }
  })
  if (comments_request.ok) {
    const comments = JSON.parse(await comments_request.json())
    const comments_element = document.getElementById("comments")
    comments.forEach((comment) => {
      const comment_element = document.createElement("div");
      comment_element.innerHTML = `
        <h3>Author Id: ${comment.Author}</h3>
        <p>${comment.Comment_content}</p>
        <p class="data">${comment.Created_at}</p>
        <hr>
        `
        comments_element.appendChild(comment_element)
    });
  }
}

async function SendCreateArticleRequest() {
  const title = document.getElementById("title").value
  const content = document.getElementById("content").value
  const turnstile_token = turnstile.getResponse()

  const create_request = await fetch("/api/articles", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Title: title,
      Content: content,
      Turnstile_token: turnstile_token
      })
  })
  if (create_request.ok) {
    const message = JSON.parse(await create_request.json())
    if (message != "Success!") {
      turnstile.reset()
      document.getElementById("status").textContent = message
    }
    document.getElementById("status").textContent = message
  }
}

async function SendDeleteArticleRequest() {
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
   window.location.replace("/")
  }
}

async function SendUpdateArticleRequest() {
  const new_title = document.getElementById("new_title").value
  const new_content = document.getElementById("new_content").value
  const author_id = document.getElementById("author_id").textContent
  const turnstile_token = turnstile.getResponse()

  const update_request = await fetch("/api/articles", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Id: Number(GetPathValue()),
      Title: new_title,
      Content: new_content,
      Author: Number(author_id),
      Turnstile_token: turnstile_token
    })
  })
  if (update_request.ok) {
    const message = JSON.parse(await update_request.json())
    if (message != "Success!") {
      turnstile.reset()
      document.getElementById("status").textContent = message
    }
    document.getElementById("status").textContent = message
 }
}

function UpdatePageRedirect(article_id) {
  window.location.replace(`/article/update/${article_id}`)
}

async function SendCreateCommentRequest() {
  const comment_content = document.getElementById("comment_content").value
  const post_id = GetPathValue()
  const turnstile_token = turnstile.getResponse()

  const create_request = await fetch("/api/comments", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Comment_content: comment_content,
      Post_id: Number(post_id),
      Turnstile_token: turnstile_token
      })
  })
  if (create_request.ok) {
    const message = JSON.parse(await create_request.json())
    if (message != "Success!") {
      turnstile.reset()
      document.getElementById("status").textContent = message
    }
    document.getElementById("status").textContent = message
  }
}
