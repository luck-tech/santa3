import { createFileRoute } from "@tanstack/react-router";
import { useRouter } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import axios from "axios";
import { UserInfo } from "@/types/login";

export const Route = createFileRoute("/github/callback")({
  component: GitHubCallback,
});

function GitHubCallback() {
  const router = useRouter();
  const searchParams = new URLSearchParams(router.state.location.search);
  const code = searchParams.get("code");

  const [loginStatus, setLoginStatus] = useState(true);

  const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    headers: {
      "Content-Type": "application/json",
    },
  });

  useEffect(() => {
    const fetchGitHubLogin = async () => {
      if (!code) router.navigate({ to: "/" });
      try {
        const response = await api.post<UserInfo>("/login/github", { code });
        console.log(response);
        setLoginStatus(false);
        router.navigate({ to: "/home" });
      } catch (error: unknown) {
        console.error("Login Error: ", error);
      }
    };

    fetchGitHubLogin();
  }, []);

  if (loginStatus) {
    return <p>isLoading...</p>;
  }
  return;
}

export default GitHubCallback;
