// Alpine counter animation helper
// Usage: x-data="counterAnim($el.dataset.count)" x-init="play()" x-text="display"
(function () {
  const fmt = new Intl.NumberFormat('en-US');
  const easeOutCubic = (t) => 1 - Math.pow(1 - t, 3);
  const parse = (v) => {
    return typeof v === 'number'
      ? Math.floor(v)
      : parseInt(String(v).replace(/,/g, ''), 10) || 0;
  };

  window.counterAnim = (target, { duration = 900, start: from = 0 } = {}) => ({
    from,
    end: parse(target),
    display: '0',

    start() {
      const reduced = matchMedia('(prefers-reduced-motion: reduce)').matches;
      const d = reduced ? 0 : duration;
      const t0 = performance.now();
      const tick = (now) => {
        const p = d ? Math.min(1, (now - t0) / d) : 1;
        const cur = Math.round(
          this.from + (this.end - this.from) * easeOutCubic(p)
        );
        this.display = fmt.format(cur);
        if (p < 1) requestAnimationFrame(tick);
      };
      requestAnimationFrame(tick);
    },
  });

  // Re-init Alpine components after HTMX swaps
  window.htmx?.onLoad?.((el) => window.Alpine?.initTree(el));
  document.addEventListener('htmx:afterSwap', (e) =>
    window.Alpine?.initTree(e.target)
  );
})();
