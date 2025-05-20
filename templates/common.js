{{define "common-js"}}

const cssMappings = {
  dark: {
    "--link-text-color": "white",
    "--link-text-hover-color": "teal",
    "--background-color": "#02071f",
    "--code-block-background-color": "#1e1b1b",
  },
  light: {
    "--link-text-color": "#02071f",
    "--link-text-hover-color": "#db4e4e",
    "--background-color": "white",
    "--code-block-background-color": "#1e1b1b",
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

document
  .getElementById("dark-mode-toggle-button")
  .addEventListener("click", toggleDarkMode);

updateColors();
{{end}}
