"use strict";

async function Fetch_Articles() {
  const response = await fetch("/api/articles", {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (response.ok) {
    const data = JSON.parse(await response.json());
    const main_element = document.getElementById("main");
    data.forEach((article) => {
      const link = document.createElement("a");
      link.href = `article/${article.Id}`;
      link.textContent = article.Title;

      const div = document.createElement("div");
      div.appendChild(link);

      main_element.appendChild(div);
    });
  }
}
