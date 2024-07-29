import React, {Suspense} from "react";

export function loadComponent(componentPath: string) {
    const LazyComponent = React.lazy(() => import(`./${componentPath}.tsx`));
    if (!LazyComponent) {
        return null
    }
    return (
        // <Suspense fallback={<div>Loading...</div>}>
        <Suspense>
            <LazyComponent/>
        </Suspense>
    )
}