import { useState, useEffect } from "react";
import { createClient } from "@supabase/supabase-js";

const supabase = createClient(
  import.meta.env.VITE_SUPABASE_URL,
  import.meta.env.VITE_SUPABASE_ANON_KEY
);

export default function App() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [session, setSession] = useState(null);

  // Listen to auth state changes (single source of truth)
  useEffect(() => {
    const {
      data: { subscription },
    } = supabase.auth.onAuthStateChange((_event, session) => {
      setSession(session);
    });

    return () => subscription.unsubscribe();
  }, []);

  // SIGN UP
  async function handleSignup(e) {
    e.preventDefault();

    const { data, error } = await supabase.auth.signUp({
      email,
      password,
      options: {
        data: { name },
      },
    });

    if (error) {
      console.error(error.message);
    } else {
      console.log("Signed up:", data.user);
    }
  }

  // SIGN IN
  async function handleSignin(e) {
    e.preventDefault();

    const { data, error } = await supabase.auth.signInWithPassword({
      email,
      password,
    });

    if (error) {
      console.error(error.message);
    } else {
      setSession(data.session);
    }
  }

  async function requestPasswordChange(e) {
    e.preventDefault();
    await supabase.auth.resetPasswordForEmail(email, {
      redirectTo: "http://localhost:5173/reset-password",
    });

    await supabase.auth.signOut()

    alert("We sent you a secure email to change your password.");
  }
  // LOG OUT
  async function handleLogout() {
    await supabase.auth.signOut();
    setSession(null);
  }

  // AUTHENTICATED UI
  if (session) {
    return (
      <div className="h-screen flex flex-col justify-center items-center gap-4">
        <h1 className="text-xl font-bold">Authenticated ✅</h1>
        <p>{session.user.email}</p>
        <button
          className="px-4 py-2 bg-red-600 text-white rounded"
          onClick={handleLogout}
        >
          Logout
        </button>
      </div>
    );
  }

  // UNAUTHENTICATED UI
  return (
    <div className="h-screen flex flex-col justify-center items-center gap-12">
      {/* SIGN UP */}
      <form
        onSubmit={handleSignup}
        className="flex flex-col gap-3 w-80 p-6 border rounded"
      >
        <h2 className="font-bold text-lg">Sign Up</h2>
        <input
          type="text"
          placeholder="Full Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button className="bg-green-600 text-white py-2 rounded">
          Sign Up
        </button>
      </form>

      {/* LOGIN */}
      <form
        onSubmit={handleSignin}
        className="flex flex-col gap-3 w-80 p-6 border rounded"
      >
        <h2 className="font-bold text-lg">Login</h2>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button className="bg-blue-600 text-white py-2 rounded">Login</button>
      </form>
      <form
        onSubmit={requestPasswordChange}
        className="flex flex-col gap-3 w-80 p-6 border rounded"
      >
        <h2 className="font-bold text-lg">Reset Password</h2>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <button className="bg-blue-600 text-white py-2 rounded">
          Send Verification Mail
        </button>
      </form>
    </div>
  );
}
