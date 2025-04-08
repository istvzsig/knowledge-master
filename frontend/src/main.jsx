import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./main.css";

import AppComponent from "./components/AppComponent";

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <AppComponent />
  </StrictMode>
);
