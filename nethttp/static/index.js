const features = [
  {
    icon: '<svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 gopher-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/></svg>',
    title: "Built-in Concurrency",
    summary:
      "Utilizes **Goroutines** (lightweight threads) and **Channels** for safe, easy, and highly efficient parallel execution.",
    details:
      "Go's concurrency model is based on CSP (Communicating Sequential Processes), enabling massive scalability for network services and distributed systems.",
  },
  {
    icon: '<svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 gopher-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.044M12 11h.01M12 21A9 9 0 003 12h.01M12 21v-8m0 0h.01"/></svg>',
    title: "Fast Compilation & Single Binary",
    summary:
      "Compiles incredibly fast, generating a single, static binary file with all dependencies included.",
    details:
      "This simplifies deployment significantly, especially in containerized environments like Docker and Kubernetes.",
  },
  {
    icon: '<svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 gopher-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7"/></svg>',
    title: "Simplicity and Readability",
    summary:
      "Features a clean, minimalist syntax with a small set of keywords, making the code easy to read and maintain.",
    details:
      "It enforces a standard formatting style (`gofmt`), which reduces cognitive load and makes collaboration easier across large teams.",
  },
  {
    icon: '<svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 gopher-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14l9-5-9-5-9 5 9 5zm0 0v5m0 0h-2m2 0h2"/></svg>',
    title: "Robust Standard Library",
    summary:
      "The 'batteries included' approach provides powerful packages for networking (HTTP), JSON, and cryptography out of the box.",
    details:
      "Minimizes the need for external dependencies, leading to more secure and stable applications.",
  },
];

const featuresGrid = document.getElementById("features-grid");
const ctaButton = document.getElementById("cta-button");
const ctaMessage = document.getElementById("cta-message");

// Function to render feature cards
function renderFeatures() {
  features.forEach((feature, index) => {
    const card = document.createElement("div");
    card.className =
      "feature-card bg-gray-50 p-6 rounded-lg border-2 border-gray-100 shadow-md";
    card.innerHTML = `
            <div class="flex items-center mb-3">
                ${feature.icon}
                <h2 class="text-xl font-semibold ml-3 text-gray-800">${feature.title}</h2>
            </div>
            <p class="text-gray-600 mb-3">${feature.summary}</p>
            <div id="details-${index}" class="text-sm text-gray-500 mt-2 p-3 bg-white rounded-lg border border-gray-200 hidden">
                ${feature.details}
            </div>
            <button class="toggle-details text-blue-600 text-sm font-medium hover:text-blue-800 transition duration-150 mt-2" data-target="details-${index}">
                Show Details
            </button>
        `;
    featuresGrid.appendChild(card);
  });
}

// Function to handle details toggle
function setupToggleListeners() {
  featuresGrid.addEventListener("click", (event) => {
    const button = event.target.closest(".toggle-details");
    if (button) {
      const targetId = button.getAttribute("data-target");
      const targetElement = document.getElementById(targetId);

      if (targetElement.classList.contains("hidden")) {
        // Show details
        targetElement.classList.remove("hidden");
        button.textContent = "Hide Details";
      } else {
        // Hide details
        targetElement.classList.add("hidden");
        button.textContent = "Show Details";
      }
    }
  });
}

// Setup CTA button click listener
ctaButton.addEventListener("click", () => {
  ctaMessage.classList.remove("hidden");
  ctaMessage.style.opacity = "0";
  setTimeout(() => {
    ctaMessage.style.opacity = "1";
  }, 10); // Small delay to trigger transition

  // In a real app, this would redirect or open a modal
  console.log("CTA clicked: Discovering Go.");
});

// Initialize the page content
window.onload = () => {
  renderFeatures();
  setupToggleListeners();
};
