import { useEffect, useState } from "react";
import { createClient } from "@supabase/supabase-js";

const supabase = createClient(
  import.meta.env.VITE_SUPABASE_URL,
  import.meta.env.VITE_SUPABASE_ANON_KEY
);

export default function ResetPassword() {
  const [password, setPassword] = useState("");
  const [allowed, setAllowed] = useState(false);
  const [loading, setLoading] = useState(true);

  // Allow password update ONLY if session comes from reset token
  useEffect(() => {
    supabase.auth.getSession().then(({ data }) => {
      if (data.session && data.session.user?.recovery_sent_at) {
        setAllowed(true);
      }
      setLoading(false);
    });
  }, []);

  async function handleUpdatePassword(e) {
    e.preventDefault();

    const { error } = await supabase.auth.updateUser({ password });

    if (error) {
      console.error(error.message);
    } else {
      alert("Password updated successfully. Please log in again.");
      await supabase.auth.signOut();
      window.location.href = "/";
    }
  }

  if (loading) return <p>Validating reset link…</p>;

  if (!allowed) {
    return <p>This password reset link is invalid or expired.</p>;
  }

  return (
    <div className="h-screen flex flex-col justify-center items-center gap-12">
      <form className="flex flex-col w-80 border rounded gap-3" onSubmit={handleUpdatePassword}>
        <h2>Set a new password</h2>
        <input
          type="password"
          placeholder="New password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button className="bg-green-600 text-white py-2 rounded border">Update password</button>
      </form>
    </div>
  );
}
