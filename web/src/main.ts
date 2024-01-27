import "./index.css"

// TODO: set tz cookie

// ---------- htmx

import "./htmx.js"

// ------------- Toastify

import Toastify from 'toastify-js'

document.body.addEventListener("toast", function(evt: any) {
  const content = document.createElement("div")
  content.textContent = evt.detail.value
  content.className = "flex-1"

  Toastify({
    node: content,
    duration: 3000,
    close: true,
    className: "bg-green-500 text-white rounded p-2 items-center gap-2 flex flex-row",
    gravity: "bottom",
    position: "center",
    stopOnFocus: true,
  }).showToast();
})

// ---------- Lit

import "./x-json.ts"
