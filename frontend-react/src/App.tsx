import { googleLogout, useGoogleLogin } from "@react-oauth/google";
import axios from "axios";

function App() {
  const login = useGoogleLogin({
    onSuccess: async (codeResponse) => {
      console.log("codeResponse", codeResponse);

      // send codeResponse to the server
      const tokenResponse = await axios.get(
        `http://localhost:8080/auth/google/callback?code=${codeResponse.code}`
      );

      console.log("tokenResponse", tokenResponse);
    },
    flow: "auth-code",
  });

  return (
    <>
      <button onClick={() => login()}>Login with google</button>

      <button onClick={() => googleLogout()}>Logout</button>
    </>
  );
}

export default App;
