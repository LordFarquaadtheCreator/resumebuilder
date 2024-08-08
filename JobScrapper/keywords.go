// contains technical keywords that show up in job descriptions
package JobScrapper

// should maybe have different arrays for different types of keywords

var KeyWords = [...]string{
    // General
    "nodejs", "node", "bun", "JS", "javascript", "typescript", "TS", 
    "web", "frontend", "backend", "fullstack", "full-stack", "full stack", 
    "developer", "development", "programming", "software", "engineer", "software engineer", 
    "web development", "software development", "full-stack developer",
    "app", "app development", "mobile app", "web app", "hybrid app",
    "native app", "cross-platform", "mobile development",
    "software engineer", "programmer", "coder", "developer",
    "it", "information technology", "technology",
    "software development lifecycle", "sdlc",
    "agile", "scrum", "kanban",
    "devops", "cloud", "aws", "azure", "gcp",
    "docker", "kubernetes",

    // Programming Languages
    "python", "java", "c++", "c#", "php",
    "ruby", "golang", "swift", "kotlin",
    "dart", "rust", "scala", "perl",
    "objective-c",

    // Frontend
    "react", "angular", "vue", "svelte", "jsx", "tsx",
    "html", "css", "scss", "less",
    "ui", "ux", "frontend developer", "front-end developer",
    "user interface", "user experience", "responsive design",
    "accessibility", "a11y", "web accessibility",
    "framework", "library", "component",
    "state management", "redux", "context", "vuex",
    "routing", "react router", "angular router", "vue router",
    "testing", "jest", "react testing library",
    "performance optimization", "seo",

    // Backend
    "express", "next", "nuxt", "sapper", "sveltekit",
    "nodejs", "node", "bun", "JS", "javascript", "typescript", "TS",
    "database", "sql", "mysql", "postgresql", "sqlite",
    "nosql", "mongodb", "redis",
    "api", "rest", "graphql",
    "serverless", "aws", "azure", "gcp",
    "cloud", "server", "deployment",
    "backend developer", "back-end developer",
    "server-side rendering", "ssr",
    "nodemon", "pm2",
    "docker", "kubernetes",
    "testing", "jest", "mocha", "chai",
    "security", "authentication", "authorization",

    // App Development
    "ios", "iphone", "ipad", "macos",
    "android", "google play",
    "flutter", "react native",
    "xamarin", "ionic",
    "mobile friendly", "mobile optimization",
    "app store optimization", "aso",
    "in-app purchase",
    "push notifications",
    "user acquisition",
    "app analytics",
    "app marketing",
}
