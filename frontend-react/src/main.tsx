import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";

import { GoogleOAuthProvider } from "@react-oauth/google";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <GoogleOAuthProvider clientId="450868207826-aaqcprl1geb0la920mkppcstu1ta3e3u.apps.googleusercontent.com">
      <App />
    </GoogleOAuthProvider>
  </React.StrictMode>
);
//
