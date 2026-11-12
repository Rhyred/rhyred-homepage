/* RHYRED — Main JavaScript */
(function () {
  'use strict';

  // ---- Mobile Navigation ----
  function initMobileNav() {
    var toggle = document.getElementById('navbar-toggle');
    var nav = document.getElementById('navbar-nav');
    if (!toggle || !nav) return;

    toggle.addEventListener('click', function () {
      var expanded = toggle.getAttribute('aria-expanded') === 'true';
      toggle.setAttribute('aria-expanded', String(!expanded));
      toggle.classList.toggle('active');
      nav.classList.toggle('active');
      document.body.style.overflow = nav.classList.contains('active') ? 'hidden' : '';
    });

    // Close nav on link click
    var links = nav.querySelectorAll('.navbar__link');
    links.forEach(function (link) {
      link.addEventListener('click', function () {
        toggle.classList.remove('active');
        nav.classList.remove('active');
        toggle.setAttribute('aria-expanded', 'false');
        document.body.style.overflow = '';
      });
    });
  }

  // ---- Navbar scroll effect ----
  function initNavbarScroll() {
    var navbar = document.getElementById('navbar');
    if (!navbar) return;

    var lastScroll = 0;
    window.addEventListener('scroll', function () {
      var currentScroll = window.pageYOffset;
      if (currentScroll > 100) {
        navbar.style.borderBottomColor = 'var(--color-border)';
      } else {
        navbar.style.borderBottomColor = 'var(--color-border-light)';
      }
      lastScroll = currentScroll;
    }, { passive: true });
  }

  // ---- Reveal on scroll ----
  function initRevealOnScroll() {
    var prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    if (prefersReducedMotion) return;

    var sections = document.querySelectorAll('.section, .page-hero');
    if (!sections.length) return;

    var observer = new IntersectionObserver(function (entries) {
      entries.forEach(function (entry) {
        if (entry.isIntersecting) {
          entry.target.classList.add('visible');
          observer.unobserve(entry.target);
        }
      });
    }, {
      threshold: 0.1,
      rootMargin: '0px 0px -40px 0px'
    });

    sections.forEach(function (section) {
      section.classList.add('reveal');
      observer.observe(section);
    });
  }

  // ---- Terminal typing effect ----
  function initTerminalEffect() {
    var prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    if (prefersReducedMotion) return;

    var outputs = document.querySelectorAll('.hero .terminal__output p');
    if (!outputs.length) return;

    outputs.forEach(function (el, i) {
      el.style.opacity = '0';
      el.style.transform = 'translateX(-8px)';
      el.style.transition = 'opacity 0.25s ease, transform 0.25s ease';
      setTimeout(function () {
        el.style.opacity = '1';
        el.style.transform = 'translateX(0)';
    }, 100 + (i * 120));
  });
}

  // ---- HTMX filter active state ----
  function initFilterActiveState() {
  function syncFilterWithUrl() {
    var params = new URLSearchParams(window.location.search);
    var category = params.get('category') || 'all';
    var filterContainer = document.getElementById('project-filter');
    if (!filterContainer) return;
    var buttons = filterContainer.querySelectorAll('.project-filter__btn');
    buttons.forEach(function (btn) {
      var href = btn.getAttribute('href') || btn.getAttribute('hx-get') || '';
      if (category === 'all' && (href.indexOf('category=all') !== -1 || href === '/projects')) {
        btn.classList.add('project-filter__btn--active');
      } else if (href.indexOf('category=' + encodeURIComponent(category)) !== -1 || href.indexOf('category=' + category) !== -1) {
        btn.classList.add('project-filter__btn--active');
      } else {
        btn.classList.remove('project-filter__btn--active');
      }
    });
  }

  document.body.addEventListener('htmx:afterSwap', function (evt) {
    if (evt.detail.target && evt.detail.target.id === 'project-grid') {
      var trigger = evt.detail.requestConfig && evt.detail.requestConfig.elt;
      if (trigger && trigger.classList.contains('project-filter__btn')) {
        var filterContainer = document.getElementById('project-filter');
        if (filterContainer) {
          var buttons = filterContainer.querySelectorAll('.project-filter__btn');
          buttons.forEach(function (btn) {
            btn.classList.remove('project-filter__btn--active');
          });
          trigger.classList.add('project-filter__btn--active');
        }
      }
    }
  });

  window.addEventListener('popstate', syncFilterWithUrl);
}

// ---- Initialize ----
document.addEventListener('DOMContentLoaded', function () {
  initMobileNav();
  initNavbarScroll();
  initRevealOnScroll();
  initTerminalEffect();
  initFilterActiveState();
});
}) ();

