variable "VERSION" {
    default = "test"
}

variable "PLATFORMS" {
    # https://hub.docker.com/hardened-images/catalog/dhi/golang/images
    default = [
        "linux/amd64",
        "linux/arm64",
    ]
}

variable "PRODUCTION" {
    default = [
        "buttery",
    ]
}

variable "TEST" {
    default = [
        "test-buttery",
    ]
}

group "production" {
    targets = PRODUCTION
}

group "test" {
    targets = TEST
}

group "all" {
    targets = concat(PRODUCTION, TEST)
}

target "buttery" {
    platforms = PLATFORMS
    tags = [
        "n4jm4/buttery:${VERSION}",
        "n4jm4/buttery",
    ]
}

target "test-buttery" {
    platforms = PLATFORMS
    tags = [
        "n4jm4/buttery:test",
    ]
}
