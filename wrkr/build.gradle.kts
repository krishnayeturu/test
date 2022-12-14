import com.github.jengelman.gradle.plugins.shadow.tasks.ShadowJar
import com.adarshr.gradle.testlogger.theme.ThemeType

plugins {
    `java-library`
    id("com.github.johnrengelman.shadow") version "7.+"
    id("com.adarshr.test-logger") version "3.+"
}

group = "com._2ndwatch"

java {
    toolchain {
        languageVersion.set(JavaLanguageVersion.of(17))
    }
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("org.slf4j:slf4j-log4j12:1.7.36")
    implementation("org.apache.kafka:kafka-streams:3.2.0")
    implementation("com.google.protobuf:protobuf-java:3.20.1")
    implementation("com.google.protobuf:protobuf-java-util:3.20.1")
    implementation("io.github.cdimascio:dotenv-java:2.2.4")

    testImplementation("org.apache.kafka:kafka-streams-test-utils:3.2.0")
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.8.1")
    testImplementation("org.mockito:mockito-core:4.6.1")

    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:5.8.1")
}

sourceSets.all {
    java.srcDirs("src/main/java", "src/generated/java")
}

tasks.getByName("build") {
    finalizedBy("shadowJar")
}

tasks.getByName<Jar>("jar") {
    enabled = false
}

tasks.withType<ShadowJar> {
    archiveFileName.set("main.jar")
    manifest {
        attributes["Main-Class"] = "com._2ndwatch.mswrkradmissionsservice.Main"
    }
}

tasks.getByName<Test>("test") {
    useJUnitPlatform()
    testlogger {
        theme = ThemeType.STANDARD
        showExceptions = true
        showStackTraces = true
        showFullStackTraces = false
        showCauses = true
        slowThreshold = 2000
        showSummary = true
        showSimpleNames = false
        showPassed = true
        showSkipped = true
        showFailed = true
        showOnlySlow = false
        showStandardStreams = false
        showPassedStandardStreams = true
        showSkippedStandardStreams = true
        showFailedStandardStreams = true
        logLevel = LogLevel.LIFECYCLE
    }
}