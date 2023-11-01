import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { login } from "./utils";

export function Callback() {
  const authContext = {
    auth: false,
  }

  const { hash } = useLocation(); // O que vem depois de # na URL
  const navigate = useNavigate();

  useEffect(() => {
    if (authContext.auth) {
      navigate("/login");
      return;
    }

    const searchParams = new URLSearchParams(hash.replace("#", ""));
    const accessToken = searchParams.get("access_token") as string;
    const idToken = searchParams.get("id_token") as string;
    const state = searchParams.get("state") as string;

    if (!accessToken || !idToken || !state) {
      navigate("/login");
    }

    login(accessToken, idToken, state);

  }, [hash, authContext, navigate]);

  return <div>Loading...</div>;
}