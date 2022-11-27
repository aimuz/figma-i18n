const dictionaryUrl = document.createElement("link");
dictionaryUrl.rel = "preload";
dictionaryUrl.as = "fetch";
dictionaryUrl.id = "dictionaryUrl";
dictionaryUrl.crossorigin = "anonymous";
dictionaryUrl.href = chrome.runtime.getURL("locales/zh-hans/dictionary.json");

document.documentElement.lang = "zh-hans";
// document.documentElement.dataset["react-helmet"]="lang"

document.addEventListener("DOMContentLoaded", fireContentLoadedEvent, false);

const locale = "ja";

function fireContentLoadedEvent() {
  console.log("DOMContentLoaded");

  document.documentElement.appendChild(dictionaryUrl);

  const scripts = document.getElementsByTagName("script");
  for (let item of scripts) {
    if (
      item.getAttribute("data-initial") != undefined &&
      !item.getAttribute("data-locale")
    ) {
      const json = JSON.parse(item.dataset.initial);
      json.INITIAL_OPTIONS.user_data.locale = locale;

      const script = document.createElement("script");
      script.removeAttribute("nonce");
      script.dataset["initial"] = JSON.stringify(json);
      script.dataset["locale"] = locale;
      script.innerText = item.innerText;

      const parent = item.parentNode;
      parent.insertBefore(script, item.nextSibling);
      return;
    }
  }
}
