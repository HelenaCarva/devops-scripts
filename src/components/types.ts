// types.ts
export interface DeploymentConfig {
  environment: string;
  region: string;
  clusterName: string;
  namespace: string;
}

export interface DockerImage {
  name: string;
  tag: string;
  repository: string;
}

export interface DeploymentScript {
  name: string;
  command: string;
  args: string[];
}

export interface DevOpsConfig {
  deploymentConfig: DeploymentConfig;
  dockerImage: DockerImage;
  deploymentScripts: DeploymentScript[];
}

export enum Environment {
  DEVELOPMENT = 'development',
  STAGING = 'staging',
  PRODUCTION = 'production',
}

export enum DeploymentStatus {
  PENDING = 'pending',
  IN_PROGRESS = 'in_progress',
  SUCCESS = 'success',
  FAILED = 'failed',
}