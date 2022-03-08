import React from "react";
import { Provider } from "react-redux";
import { Routes, Route } from "react-router-dom";
import Box from "@mui/material/Box";
import logoUrl from "./assets/logo.svg";
import { AssignHostname } from "./components/wizard/AssignHostname";
import { ConfigureWifi } from "./components/wizard/ConfigureWifi/ConfigureWifi";
import { Home } from "./components/wizard/Home";
import { Landing } from "./components/wizard/Landing";
import { RegisterFactoryName } from "./components/wizard/RegisterFactoryName";
import { Shell } from "./components/wizard/Shell";
import { store } from "./store";

function App() {
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        height: "100%",
        backgroundColor: "#2F2F2F",
      }}
    >
      <Provider store={store}>
        <Box
          component="header"
          sx={{ display: "flex", width: "100%", padding: "30px" }}
        >
          <img src={logoUrl} alt="Arduino Pro" />
        </Box>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/landing" element={<Landing />} />
          <Route path="/wlan" element={<ConfigureWifi />} />
          <Route path="/factory" element={<RegisterFactoryName />} />
          <Route path="/shell" element={<Shell />} />
          <Route path="/hostname" element={<AssignHostname />} />
        </Routes>
      </Provider>
    </Box>
  );
}

export default App;
