document
  .getElementById("shorten-form")
  .addEventListener("submit", async function (e) {
    e.preventDefault();

    let url = document.getElementById("url-input").value;

    let response = await fetch("/shorten", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: "url=" + encodeURIComponent(url),
    });
    let result = await response.text();
    document.getElementById("output-url").innerHTML = `
      <a href=${result}>Shortened URL: ${result}</a>
    `;
  });
