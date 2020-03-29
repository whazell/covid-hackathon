import { useEffect } from 'react';

/**
 * Util func so you can do useEffectAsync(async () => { await stuff()})
 * Does not support an unsubscriber!
 * @param effect
 * @param dependencies
 */
const useEffectAsync = (
  effect,
  dependencies
) => {
  useEffect(() => {
    effect();
  }, dependencies);
};

export default useEffectAsync;
