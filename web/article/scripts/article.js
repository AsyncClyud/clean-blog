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
    const data = JSON.parse(await response.json());
    const main_element = document.getElementById("main");
    data.forEach((article) => {
      const article_element = document.createElement("div");
      article_element.setAttribute("class", "article");
      article_element.setAttribute("id", article.Id);
      article_element.innerHTML = `
      <h3 id = "title">${article.Title}<h3>
      <p id ="content">${article.Content}</p>
      <p id="created_at">${article.Created_at}</p>
      `;
      main_element.appendChild(article_element);
      document.title = article.Title
    });
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
