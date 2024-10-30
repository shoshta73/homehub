import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

export type State = {
  stats: {
    activeView: "pastebin";
  };
};

export interface StateAction {
  toggleView: (view: "pastebin") => void;
}

const useHVState = create<State & StateAction>()(
  persist(
    (set) => ({
      stats: {
        activeView: "pastebin",
      },
      toggleView: (view: "pastebin") =>
        set((state) => {
          if (state.stats.activeView === view) {
            return state;
          }
          return { ...state, stats: { ...state.stats, activeView: view } };
        }),
    }),
    {
      name: "homeview",
      storage: createJSONStorage(() => localStorage),
    }
  )
);

export default useHVState;
