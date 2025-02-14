import { useEffect, useState } from "react";
import { doFetch, Method } from "../service/doFetch";
import { useStore } from "../store";

export interface BotConfig {
    OwnerTwitchId: string;
    JoinBot: boolean;
}

export type SetBotConfig = (config: BotConfig) => void;

export function useBotConfig(): [BotConfig | undefined, SetBotConfig, boolean] {
    const [botConfig, setBotConfig] = useState<BotConfig | undefined>(undefined);
    const [loading, setLoading] = useState(true);
    const managing = useStore(state => state.managing);
    const scToken = useStore(state => state.scToken);

    const fetchConfig = () => {
        setLoading(true);

        const endPoint = "/api/botconfig";
        doFetch(Method.GET, endPoint).then(setBotConfig).then(() => setLoading(false)).catch(() => setLoading(false));
    };

    // eslint-disable-next-line react-hooks/exhaustive-deps
    useEffect(fetchConfig, [managing, scToken]);

    const setCfg = (config: BotConfig) => {
        setLoading(true);
        
        doFetch(Method.POST, "/api/botconfig", undefined, config).then(fetchConfig).catch(() => setLoading(false));
    };

    return [botConfig, setCfg, loading]
}