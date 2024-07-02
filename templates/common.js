{{define "common-js"}}

const cssMappings = {
  dark: {
    "--link-text-color": "grey",
    "--background-color": "black",
  },
  light: {
    "--link-text-color": "#02071f",
    "--background-color": "white",
  },
};

const root = document.querySelector(":root");

function updateColors() {
  if (JSON.parse(window.localStorage.getItem("dark-mode-enabled"))) {
    Object.entries(cssMappings["dark"]).forEach(([key, value]) => {
      document.body.style.setProperty(key, value);
    });
    return;
  }
  Object.entries(cssMappings["light"]).forEach(([key, value]) => {
    document.body.style.setProperty(key, value);
  });
}

function toggleDarkMode(event) {
  const darkModeEnabled = window.localStorage.getItem("dark-mode-enabled");
  if (darkModeEnabled === "null") {
    window.localStorage.setItem("dark-mode-enabled", "true");
    return;
  }
  window.localStorage.setItem(
    "dark-mode-enabled",
    !JSON.parse(window.localStorage.getItem("dark-mode-enabled")),
  );
  updateColors();
}

updateColors();

document
  .getElementById("dark-mode-toggle-button")
  .addEventListener("click", toggleDarkMode);
{{end}}
