import Menubar from "@/components/menubar";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Pagination, PaginationContent, PaginationItem, PaginationLink } from "@/components/ui/pagination";
import useHVState from "@/store/homeviewstate";
import { ClipboardIcon } from "lucide-react";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const PastebinStats = () => {
  const [stats, setStats] = useState<{
    id: number;
    created: number;
    sharedWithMe: number;
  } | null>(null);
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    const validate = async () => {
      try {
        const res = await fetch("/auth/validate", {
          method: "POST",
          credentials: "include",
        });
        if (res.status !== 200) {
          navigate("/login");
        }
      } catch (err) {
        console.error(err);
        navigate("/login");
      }
    };

    validate();
  }, []);

  useEffect(() => {
    const getStats = async () => {
      try {
        setLoading(true);
        const res = await fetch("/pastebin/stats");
        const data = (await res.json()) as {
          id: number;
          created: number;
          sharedWithMe: number;
        };
        setStats(data);
      } finally {
        setLoading(false);
      }
    };

    if (stats === null) {
      getStats();
    }
  }, [stats, loading]);

  return (
    <>
      {loading && <div>Loading...</div>}
      {!loading && !stats && <></>}
      {!loading && stats && (
        <div>
          <div className="flex flex-row">
            <h1>created:</h1>
            <p>{stats.created}</p>
          </div>
          <div className="flex flex-row">
            <h1>sharedWithMe:</h1>
            <p>{stats.created}</p>
          </div>
        </div>
      )}
    </>
  );
};

function UserHomeView() {
  const state = useHVState();

  return (
    <>
      <Menubar />
      <div className="flex flex-col flex-grow w-full">
        <div className="flex flex-col w-2/3 mx-auto mt-4">
          <Card>
            <CardHeader>
              <CardTitle>Stats</CardTitle>
              <CardDescription>View of your stats</CardDescription>
            </CardHeader>
            <CardContent>
              <Pagination>
                <PaginationContent>
                  <PaginationItem>
                    <PaginationLink
                      isActive={state.stats.activeView === "pastebin"}
                      onClick={() => state.toggleView("pastebin")}
                    >
                      <ClipboardIcon />
                    </PaginationLink>
                  </PaginationItem>
                </PaginationContent>
              </Pagination>
              {state.stats.activeView === "pastebin" && <PastebinStats />}
            </CardContent>
          </Card>
        </div>
      </div>
    </>
  );
}

export default UserHomeView;
