import React, { useEffect, useState } from "react"

type t_RunFetch = (silent?: boolean) => void

export function useFetch<T = unknown>(key: string, dataProvider: (() => Promise<T>)): [T | undefined, boolean, unknown, t_RunFetch] {
    let [data, setData] = useState<T>();
    let [loading, setLoading] = useState<boolean>(false);
    let [error, setError] = useState<unknown>();

    const runFetch: t_RunFetch = async function (silent?: boolean) {
        setError(undefined)
        // Silent==true skips updating loading flag
        if (silent !== true) setLoading(true);
        try {
            setData(await dataProvider());
        } catch (err) {
            setError(err);
        }
        setLoading(false);
    }

    useEffect(() => {
        runFetch();
    }, []);

    return [data, loading, error, runFetch];
}

type t_ItemRenderer<T> = (arg0: T) => React.ReactNode;
type t_LoadingRenderer = () => React.ReactNode;
type t_ErrorRenderer = (arg0: unknown, arg1: t_RunFetch) => React.ReactNode;

/**
 * 
 * @param key Cache key
 * @param dataProvider Data provider function (ex. axio.get(), fetch() ...)
 * @param itemRenderer Function that provides a jsx template for each item in data
 * @param loadingRenderer Function that provides a jsx template for when data is loading
 * @param errorRenderer Function that provides a jsx template for when an error occured while loading data
 * @returns Either the rendered list of items, loading jsx or error jsx
 */
export function useTemplatedFetch<T = unknown>(
    key: string, 
    dataProvider: (() => Promise<[T]>),
    itemRenderer: t_ItemRenderer<T>,
    loadingRenderer: t_LoadingRenderer,
    errorRenderer: t_ErrorRenderer
): React.ReactNode{

    let [data, loading, error, runFetch] = useFetch(key, dataProvider)

    if (loading) {
        return loadingRenderer()
    }
    else if (typeof error !== "undefined") {
        return errorRenderer(error, runFetch)
    }

    return data?.map((value) => {
        return itemRenderer(value);
    })
}